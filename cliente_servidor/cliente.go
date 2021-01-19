/*****************************************************************************
 * cliente.go                                                                 
 * Nome:
 * Matrícula:
 *****************************************************************************/

package main

import (
  "os"
  "log"
)

const SEND_BUFFER_SIZE = 2048

/* TODO: client()
 * Abrir socket e enviar mensagem de stdin.
*/
func client(server_ip string, server_port string) {

}

// Main obtém argumentos da linha de comando e chama função client
func main() {
  if len(os.Args) != 3 {
    log.Fatal("Uso: ./cliente [IP servidor] [porta servidor] < [arquivo mensagem]")
  }
  server_ip := os.Args[1]
  server_port := os.Args[2]
  client(server_ip, server_port)
}
