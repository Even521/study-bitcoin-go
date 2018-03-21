package block

import (
	"encoding/gob"
	"log"
	"crypto/sha256"
	"bytes"
	"fmt"
	"encoding/hex"
)

const subsidy  = 10 //初始化补助为10


//交易事物
type Transaction struct {
	ID   []byte     //交易hash
	Vin  []TXInput  //事物输入
	Vout []TXOutput //事物输出
}

//一个事物输入
type TXInput struct {
	Txid []byte //交易ID的hash
	Vout int    //交易输出
	ScriptSig string //解锁脚本
}
//一个事物输出
type TXOutput struct {
	Value int       //值
	ScriptPubKey string //解锁脚本key
}



func (tx *Transaction)IsCoinbase()   bool  {
    return len(tx.Vin)==1&&len(tx.Vin[0].Txid)==0&&tx.Vin[0].Vout==-1
}
//设置交易ID hash
func (tx *Transaction) SetID(){
	var encoded bytes.Buffer
	var hash [32]byte //32位的hash字节

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	//将交易信息sha256
	hash = sha256.Sum256(encoded.Bytes())
	//生成hash
	tx.ID = hash[:]
}

//创建Coinbase事物
func NewCoinbaseTX(to,data string) *Transaction  {
	if data==""{
		data=fmt.Sprintf("Reward to '%s'",to)
	}
    //这里Vout-1 data：const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"
	txin :=TXInput{[]byte{},-1,data}
	//subsidy是奖励的金额
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	//设置32位交易hash
	tx.SetID()
	return &tx
}

//通过检查地址是否启动了事务
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}
//检查输出是否可以使用所提供的数据进行解锁
func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}
//创建一个新的未经使用的交易输出
func NewUTXOTransaction(from, to string, amount int, bc *Blockchain)   *Transaction{
	var inputs []TXInput
	var outputs []TXOutput
    //查询发币地址所未经使用的交易输出
	acc, validOutputs := bc.FindSpendableOutputs(from, amount)
    //判断是否有那么多可花费的币
	if acc < amount {
		log.Panic("ERROR: Not enough funds")
	}
	// Build a list of inputs
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			log.Panic(err)
		}
		for _, out := range outs {
			input := TXInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}
	// Build a list of outputs
	outputs = append(outputs, TXOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, TXOutput{acc - amount, from}) // a change
	}
	tx := Transaction{nil, inputs, outputs}
	tx.SetID()
	return &tx
}