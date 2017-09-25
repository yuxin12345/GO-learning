package main
import (
	"net/rpc"
	"fmt"
	"net/http"
	"encoding/gob"
	"bufio"
	"time"
	"net"
	"io"
)
type gobClientCodec1 struct{
	rwc io.ReadWriteCloser
	dec *gob.Decoder
	enc *gob.Encoder
	encBuf *bufio.Writer
} 
func (c *gobClientCodec1) WriteRequest(r *rpc.Request,boy interface{})(err error){
	if err =TimeoutCoder(c.enc.Encode,r,"client write request"):err!=nil{
		return
	}
	return c.encBuf.Flush()
}
func (c *gobClientCodec1) ReadResponseHeader(r *rpc.Response)error{
	return c.dec.Decode(r)
}
func (c *gobClientCodec1) ReadResponseBody(body interface{})error{
	return c.dec.Decode(body)
}
func (c *gobClientCodec1) Close()error{
	return c.rwc.Close()
}
func call(srv string,rpcname string,age interface{},reply interface{})error{
	conn,err :=net.DialTimeout("tcp",srv+":4200",time.Second*10)
	if err != nil{
		return fmt.Errorf("ConnectError:%s",err.Error())
	}
	encBuf :=bufio.NewWriter(conn)
	codec := &gobClientCodec1
}