/*****************************************************************************
 * cliente.go
 * Nome: Henrique Lopes Lima, Gabriela Miranda Leal
 * Matrícula: 413031, 398624
 *****************************************************************************/

package main

import (
  "bufio"
  "fmt"
  "io"
  "log"
  "net"
  "os"
)

const SendBufferSize = 2048

/* TODO: client()
 * Abrir socket e enviar mensagem de stdin.
*/
func client(serverIp string, serverPort string) {
  //TCPAddr
  tcpAddr, err := net.ResolveTCPAddr("tcp", serverIp + serverPort)
  checkErrorClient(err)

  //TCPConn
  conn, err := net.DialTCP("tcp", nil, tcpAddr)
  checkErrorClient(err)
  reader := bufio.NewReader(os.Stdin)
  buf := make([]byte, SendBufferSize)

  for {
    readTotal, err := reader.Read(buf)
    if err != nil {
      if err != io.EOF {
        checkErrorClient(err)
      }
      break
    }
    _, err = conn.Write(buf[:readTotal])
    checkErrorClient(err)
  }

  checkErrorClient(err)
  os.Exit(0)
}

// Main obtém argumentos da linha de comando e chama função client
func main() {
  if len(os.Args) != 3 {
    log.Fatal("Uso: ./cliente [IP servidor] [porta servidor] < [arquivo mensagem]")
  }
  serverIp := os.Args[1]
  serverPort := os.Args[2]
  client(serverIp, serverPort)
}

func checkErrorClient(err error){
  if err != nil {
    _, _ = fmt.Fprintf(os.Stderr, "Erro: %s\n", err.Error())
    os.Exit(1)
  }
}
