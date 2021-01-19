/*****************************************************************************
 * servidor.go                                                                 
 * Nome:
 * Matrícula:
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

  for {
    conn, err := listener.Accept()
    fmt.Println("Conectado")

    if err != nil {
      continue
    }

    go handleClient(conn)
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

func handleClient(conn net.Conn)  {
  defer conn.Close()

  var buf [RecvBufferSize]byte

  for {
    n, err := conn.Read(buf[0:])
    if err != nil {
      return
    }

    fmt.Println(string(buf[0:]))

    _, err2 := conn.Write(buf[0:n])

    if err2 != nil {
      return
    }
  }
}
