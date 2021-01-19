/*****************************************************************************
 * servidor.go                                                                 
 * Nome:
 * Matrícula:
 *****************************************************************************/

package main

import (
  "os"
  "log"
)

const RECV_BUFFER_SIZE = 2048

/* TODO: server()
 * Abra socket e espere o cliente conectar
 * Imprima a mensagem recebida em stdout
*/
func server(server_port string) {

}


// Main obtém argumentos da linha de comando e chama a função servidor
func main() {
  if len(os.Args) != 2 {
    log.Fatal("Uso: ./servidor [porta servidor]")
  }
  server_port := os.Args[1]
  server(server_port)
}
