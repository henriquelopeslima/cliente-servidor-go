/*****************************************************************************
 * cliente.go                                                                 
 * Nome:
 * Matrícula:
 *****************************************************************************/

package main

import (
  "bufio"
  "fmt"
  "log"
  "net"
  "os"
  "strings"
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

  message := readInput()
  _, err = conn.Write([]byte(message))
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
func readInput() string {
  texts := make([]string, 0)
  scanner := bufio.NewScanner(os.Stdin)
  for {
    scanner.Scan()
    text := scanner.Text()
    if len(text) != 0 {
      texts = append(texts, text)
    } else {
      break
    }
  }
  stringText := strings.Join(texts[:], "\n")
  return stringText
}

func checkErrorClient(err error){
  if err != nil {
    _, _ = fmt.Fprintf(os.Stderr, "Erro: %s\n", err.Error())
    os.Exit(1)
  }
}
