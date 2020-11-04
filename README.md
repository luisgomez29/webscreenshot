# Website Screenshot (golang)

Microservicio que toma capturas de pantalla de un sitio web dada su URL y devuelve la URL de la imagen.

## Dependecias

- [gorilla/mux](https://github.com/gorilla/mux) Router HTTP
- [chromedp](https://github.com/chromedp/chromedp) Una forma más rápida y sencilla de impulsar navegadores compatibles con
el protocolo Chrome DevTools.

## Instalación

1. Clonar el proyecto.

4. Ejecutar el servidor local:
```bash
# Ejecutar servidor
go run main.go
```

## Uso

Ingresar en el navegador a la ruta `/screenshot` luego pasar el parametro `url` con la URL del sitio web a tomar el 
pantallazo: 

    localhost:8000/screenshot?url=https://www.google.com
o también:

    localhost:8000/screenshot?url=https://google.com

Las capturas de pantalla se almacenan en la carpeta `images`.

## Autor

**Luis Guillermo Gómez**  
- [Github](https://github.com/luisgomez29)

`Gracias!.`