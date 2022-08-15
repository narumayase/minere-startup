# minere-startup

Aplicación para manejar la minería de eth.

## casos de uso

- cuando comienza envía un email que sirve como alerta para saber que la computadora ha iniciado
- si el archivo [db.data](https://github.com/narumayase/minere-startup/blob/main/db.data) fue escrito con un número menor al umbral de tolerancia configurado (ver en [configuración](https://github.com/narumayase/minere-startup#confjson)), la minería comienza, sino, envía un mail que sirve de alerta de que algo ha fallado muchas veces y requiere revisión manual.

## build
Instalar [Golang](https://golang.org/) para construir esta aplicación y ejecutar el comando debajo.

```
$ go build 
```

## pre-requisitos

* copiar el ejecutable creado por ```go build``` en la carpeta de inicio del sistema operativo o configurar el sistema operativo para que inicie esta aplicación cuando inicia
* por ahora solo funciona con teamredminer con estos argumentos:
    * ```-a ethash -o pool -u ****.trmtest -p x```
    * copiar el archivo [conf.json](https://github.com/narumayase/minere-startup/blob/main/examples/conf.json) al lado del ejecutable de esta aplicación y configúrelo

### conf.json

| parámetros    |  descripción
|---------------|--------------------------------------------------
| **Threshold** | umbral de tolerancia para el reinicio del sistema
| **Email**     | ver la descripción de Email debajo
| **Miner**     | ver la descripción de Miner debajo

| Email parámetros |  descripción
|------------------|--------------------------------------------------
| **From**         | dirección de email desde donde esta aplicación envía el mail
| **To**           | dirección de email a la que esta aplicación envía el mail
| **Smtp**         | servidor smtp
| **SmtpPort**     | puerto smtp
| **PasswordFrom** | contraseña del email desde el que esta aplicación envía el mail
| **Start**        | ver la descripción de Start debajo
| **Stop**         | ver la descripción de Stop debajo

| Start/Stop parámetros |  descripción
|-----------------------|--------------------------------------------------
| **Body**              | cuerpo de mensaje para enviar en el mail, puede ser en texto plano e incluso en formato html
| **Subject**           | asunto del email a enviar
| **Send**              | si está habilitado a enviar este email. puedes configurar esta aplicación para que no envíe email en el inicio o en el stop de la aplicación

| Miner parámetros |  descripción
|------------------|--------------------------------------------------
| **Path**         | ruta del archivo ejecutable donde se encuentra el teamredminer
| **User**         | dirección de billetera del usuario + .trmtest
| **Pool**         | pool de minería

* example
```
{
  "Email": {
    "Start": {
      "Body": "Hi! I'm Ms. Minere Traba Jadora, I work for you, do you remember? Well.. I had some issues, I don't know what happened but you know, I just crashed.\nBut hey! I'm on again and I will try to start the mining again.\n\nBest regards.",
      "Subject": "ALERT - Minere needs your help!",
      "Send": true
    },
    "Stop": {
      "Body": "Errmm... It's me... again. \nI restarted the system two times and I'm afraid to continue.. Can you tell me what to do?\nI'm sorry",
      "Subject": "ALERT - Minere REALLY needs your help!",
      "Send": true
    },
    "From": "****@gmail.com",
    "To": "****@gmail.com",
    "Smtp": "smtp.gmail.com:587",
    "SmtpPort": 587,
    "PasswordFrom": "****"
  },
  "Miner": {
    "Path": "C:\\Users\\Default\\Documents\\teamredminer-v0.8.5-win\\teamredminer.exe",
    "User": "*****.trmtest",
    "Pool": "stratum+tcp://us1.ethermine.org:4444"
  },
  "Threshold": 2
}
```
