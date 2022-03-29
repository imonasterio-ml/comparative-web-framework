# Comparative web framework

## Los competidores

- `Gin`: Uno de los frameworks web más populares, se basa en un enfoque basado en enrutador y middleware, donde se define los controladores (handlres) adjuntos a rutas específicas. Debido a su fuerte enfoque en la extensibilidad, puede crear una solución adecuada para casi todos los escenarios de aplicaciones web.  Incluye características útiles como query params,  manejo de JSON a structs, middleware, entre otros.  Contiene casi todo lo que necesita para la mayoría de los casos de uso
- `Echo`:  Expone una estructura sorprendentemente similar en comparación con gin, proporcionando una gran cantidad de funciones de middleware disponibles como soporte de autenticación JWT, manejo de CORS, etc. Una gran diferencia es la lógica de enrutamiento: mientras que gin se basa en la biblioteca fasthttprouter (investigar) con todos sus ventajas y desventajas, echo maneja su enrutamiento utilizando un módulo autónomo.
- `Chi`:  Es una de las soluciones básicas que se recomienda usar para proyectos en los que desea obtener los beneficios de una excelente lógica de manejo de rutas, pero también se prefiere a las API de bajo nivel para controlar todo el ciclo de vida de la solicitud (del request). Al igual que con las otras librerías, las funcionalidades de middleware también son muy amigables, por lo que no está obligado a crearlas por uno mismo. Chi funciona muy bien en combinación con la librería mux.

## Comparativa de sintaxis

* [Gin](/gin/main.go)
* [Echo](/echo/main.go)
* [Chi](/chi/)

### Un simple GET

Mirando el código respectivo al Get path y su respectivo controlador `UserGetHandler`, se puede ver el alcance de las funciones expuestas por cada librería: mientras que echo y gin usan su propio context y brindan ayuda para dar formato a las respuestas, chi utiliza controladores estándar para procesar los request. Lo que significa que tendría que usar otras dependencias o crear sus propias funcionalidades en torno a esto, como hice en el ejemplo. Sin embargo, todas las librerías incluyen utilidades auxiliares básicas, como la extracción de parámetros del path

### Analizando un POST

Observando ahora el código del path respectivo al Post con su controlador `UserPostHandler`, una vez más podemos ver las diferencias entre las librerías minimalistas (Chi) y las librerías que tienen todas las funciones. Especialmente para vincular el body del request con las estructuras de Go (JSON -> struct). Es muy trabajoso codear la funcionalidad de “leer” el body del request y desarmarlo hacia la estructura (Marshaling) a diferencia de usar una única línea de código como es el caso de gin y echo.

### Conclusión

Al usar bibliotecas completas como gin y echo, se obtienen muchos beneficios para escribir funcionalidad en menos líneas de código y, a cambio, sacrificar algo de flexibilidad y, en algunos casos, rendimiento. Elegir una biblioteca auxiliar de enrutamiento muy liviana como chi puede ayudarlnos a crear aplicaciones web realmente rápidas, pero debemos dedicar más tiempo a implementar o buscar la funcionalidad que falta para analizar los cuerpos de las solicitudes, manejar las respuestas, etc.

En mi opinión, el común de los proyectos no requieren ese plus de velocidad que ofrece Chi. Por lo que preferiría utilizar librerías que brinden una simpleza en su código. En especial el punto donde ganan Gin y Echo es en el `manejo de errores`.
Por eso, creo que es mejor analizar con mas detalle a Gin y Echo en el siguiente [enlace](https://mattermost.com/blog/choosing-a-go-framework-gin-vs-echo/). Pero a modo resumen  agrego el cuadro comparativo del mismo enlace.

|      | Gin                                                                                                                                    | Echo                                                                                                                                       |
| ---- | -------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------ |
| Pros | \- High performance<br>\- Crash-free<br>\- Great error managment<br>\- Easy JSON validation<br>\- Room for control and experimentation | \- Great template support<br>\- Highly-scalable<br>\- Automatic TLS<br>\- Certifcates<br>\- Broad file type support<br>\- Easly extensible |
| Cont | \- Less documentation<br>\- Less-concise syntax<br>\- Not as flexible in development                                                   | \- Slower performance<br>\- Smaller community for support<br>\- Less control over minute details                                           |


