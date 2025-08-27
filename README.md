1. Colocar las librerias/dependencias en el archivo go.mod

#  Comandos para levantar el servicio
1. go mod tidy (instala librerias)
2. (opcional) caching de librerias:  go mod download
3. go build -o server ./cmd/server (crear el binario)
4. ./server
