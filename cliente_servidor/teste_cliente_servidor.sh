#!/bin/bash
##
## SYNOPSE
##    teste_cliente_servidor
##
## DESCRIÇÃO
##    Script de teste para o trabalho 1 de rus0082.
##    Executa 5 testes diferentes da comunicação cliente/servidor.

# Verifica o número correto de argumentos
if [[ $# -ne 1 ]]; then
  printf "USO: $0 [porta servidor]\n"
  exit
fi

WORKSPACE=/vagrant/trabalho1/.workspace
numCorrect=0
TESTS_PER_IMPL=1 # LEMBRE-SE DE ATUALIZAR SE A QTD DE TESTES MUDAR!
PORT=$1
SKIP_MESSAGE="Um ou ambos programas faltando. Pulando. \n\n"
testNum=1

# Localização dos arquivos
SGC=/vagrant/trabalho1/cliente_servidor/cliente # cliente
SGS=/vagrant/trabalho1/cliente_servidor/servidor # servidor

# função para comparar os arquivos de mensagem
# $1 = primeiro arquivo, $2 = segundo arquivo, $3 = imprime separador (não se 0, sim caso contrário),
# $4 = imprime diff (não se 0, sim caso contrário)
function compare {
  if diff -q $1 $2 > /dev/null; then
    printf "\nSUCESSO: Mensagem recebida corresponde à mensagem enviada!\n"      
    ((numCorrect++))
  else
    printf "\nFALHA: Mensagem recebida não corresponde à mensagem enviada.\n"
    if [ $4 -ne 0 ]; then
      echo Diferenças:
      diff $1 $2
    fi
  fi
  if [ $3 -ne 0 ]; then
    printf "________________________________________"               
  fi
  printf "\n" 
}

# $1 = cliente, $2 = servidor, $3 = porta, $4 = imprime separador (não se 0, sim caso contrário),
# $5 = imprime diff (não se 0, sim caso contrário)
function test {
  $2 $3 > test_output.txt &
  SERVER_PID=$!
  sleep 1
  $1 127.0.0.1 $3 < test_message.txt >/dev/null
  EXIT_STATUS=$?
  sleep 0.2
  kill $SERVER_PID
  wait $SERVER_PID 2> /dev/null
  sleep 0.2
  compare test_message.txt test_output.txt $4 $5
  rm -f test_output.txt
  sleep 0.2
}

#####################################################
# Testes, $1 = cliente, $2 = servidor, $3 = porta
#####################################################

function all-tests {

  printf "\n$testNum. TESTA MENSAGEM PEQUENA\n"
  printf "Vai filhão!\n" > test_message.txt
  test "$1" "$2" $3 1 1
  ((testNum++))

  ###############################################################################

  printf "\n$testNum. TESTA MENSAGEM ALFANUMÉRICA ALEATÓRIA\n"
  head -c100000 /dev/urandom | LC_ALL=C tr -dc 'a-zA-Z0-9' > test_message.txt
  test "$1" "$2" $3 1 0
  ((testNum++))

  ###############################################################################

  printf "\n$testNum. TESTA MENSAGEM BINÁRIA ALEATÓRIA\n"                                             
  head -c100000 /dev/urandom > test_message.txt
  test "$1" "$2" $3 1 0
  ((testNum++))

  ###############################################################################

  printf "\n$testNum. TESTA O LAÇO INFINITO DO SERVIDOR (multiplos clientes sequenciais no mesmo servidor)\n" 
  $2 $3 > test_output.txt &
  SERVER_PID=$!
  sleep 0.2
  printf "Line 1\n" | $1 127.0.0.1 $3 >/dev/null
  sleep 0.1
  printf "Line 2\n" | $1 127.0.0.1 $3 >/dev/null
  sleep 0.1
  printf "Line 3\n" | $1 127.0.0.1 $3 >/dev/null
  sleep 0.1
  kill $SERVER_PID
  wait $SERVER_PID 2> /dev/null
  sleep 0.2
  printf "Line 1\nLine 2\nLine 3\n" > test_message.txt
  compare test_message.txt test_output.txt 1 1
  rm -f test_output.txt
  sleep 0.2
  ((testNum++))

  ###############################################################################

  # Envia 10 mensagens alfanuméricas aleatórias ao servidor concorrentemente.
  # Desde que a ordem de envio e de recepção pode diferir, as mensagens são ordenadas
  # em cada ponto da comunicação para verificar a correspondência.

  printf "\n$testNum. TESTA FILA DO SERVIDOR (sobrepõe clientes no mesmo servidor)\n"                 
  rm -f test_message.txt
  stdbuf -i0 -o0 $2 $3 > test_output.txt &
  SERVER_PID=$!
  sleep 0.2
  for i in {0..9}; do
  	(timeout 1 cat /dev/urandom | LC_ALL=C tr -dc 'a-zA-Z0-9' ; echo) | tee test_message$i.txt | $1 127.0.0.1 $3 >/dev/null &
    CLIENT_PID[$i]=$!
  done
  sleep 2
  for i in {0..9}; do
  	cat test_message$i.txt >> test_message.txt
  done
  sort test_message.txt > test_message_sorted.txt
  rm -f test_message.txt
  kill $SERVER_PID
  wait $SERVER_PID 2> /dev/null
  sleep 1
  sort test_output.txt > test_output_sorted.txt
  rm -f test_output.txt
  compare test_message_sorted.txt test_output_sorted.txt 0 0
  rm -f test_output_sorted.txt test_message_sorted.txt
  sleep 0.2
  ((testNum++))
}

function handle_interrupt {
  kill $SERVER_PID 2> /dev/null
  wait $SERVER_PID &> /dev/null
  for i in {0..9}; do
    kill ${CLIENT_PID[$i]} 2> /dev/null
    wait ${CLIENT_PID[$i]} &> /dev/null
  done
  rm -rf $WORKSPACE
  echo ""
  exit 1
}
# mata o servidor no caso do sinal SIGINT
trap handle_interrupt SIGINT

####################################################
# EXECUTA TESTES
####################################################

rm -rf $WORKSPACE
mkdir $WORKSPACE
cd $WORKSPACE

if [[ -f $SGC && -f $SGS ]]; then
    all-tests $SGC $SGS $PORT
else
    printf "\n$SKIP_MESSAGE"
    ((testNum+=$TESTS_PER_IMPL))
fi

rm -rf $WORKSPACE

#####################################################
# RESUMO DOS RESULTADOS
#####################################################

printf "================================================================\n\n"
printf "TESTES QUE PASSARAM: $numCorrect/$((testNum-1))\n"
