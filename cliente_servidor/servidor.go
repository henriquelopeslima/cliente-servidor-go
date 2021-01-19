/*****************************************************************************
 * servidor.go                                                                 
 * Nome: Henrique Lopes Lima
 * Matrícula: 413031
 *****************************************************************************/

package main

import (
  "fmt"
  "log"
  "net"
  "os"
)

const RecvBufferSize = 2048

/* TODO: server()
 * Abra socket e espere o cliente conectar
 * Imprima a mensagem recebida em stdout
*/
func server(serverPort string) {

  tcpAddr, err := net.ResolveTCPAddr("tcp", serverPort)
  checkErrorServer(err)

  listener, err := net.ListenTCP("tcp", tcpAddr)
  checkErrorServer(err)

  defer listener.Close()

  for {
    conn, err := listener.Accept()

    if err != nil {
      continue
    }

    handleClient(conn)
  }
}

func handleClient(conn net.Conn)  {
  var buf [RecvBufferSize]byte

  for {
    n, err := conn.Read(buf[0:])
    if err != nil {
      return
    }
    if buf[0:] != nil {
      fmt.Print(string(buf[0:n]))
    }
  }
}

// Main obtém argumentos da linha de comando e chama a função servidor
func main() {
  if len(os.Args) != 2 {
    log.Fatal("Uso: ./servidor [porta servidor]")
  }
  serverPort := os.Args[1]
  server(serverPort)
}

func checkErrorServer(err error){
  if err != nil {
    _, _ = fmt.Fprintf(os.Stderr, "Erro: %s\n", err.Error())
    os.Exit(1)
  }
}