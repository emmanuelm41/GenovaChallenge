# GENOVA CHALLENGE

**Descripcion**

En esta ocasion el challenge se encuentra organizado en 3 etapas. Cada una de ellas se apoya sobre la anterior. Para la primera de ellas tomamos los siguientes lineamientos:
- Para el calculo de la trilateración utilizaremos una libreria externa disponible llamada github.com/co60ca/trilateration. Esta nos permitira desligarnos de los calculos matematicoa involucrados.
- Para la decodificacion del mensaje recibido utilizaremos un algoritmo propio.

Toda la logica necesaria para ambas tareas estara contenida en un mismo modulo. Alrededor de eso construiremos el modulo HTTP correspondiente para poder exponerlas al mundo exterior. Esta es la segunda etapa de este challenge. El modulo HTTP estara conenido en un modulo particular, el cual implementara los handlers necesarios para cada una de las rutas requeridas. 


**Start Locally**

To exec app run "go run ." on working dir of the project


**Use Public URL (Google App Engine)**
- TopSecret service --> http://genovachallenge.uc.r.appspot.com:80/topsecret
- TopSecretSplit service --> http://genovachallenge.uc.r.appspot.com:80/topsecret_split

Use some of the examples inside the folder "examples" to test the app. They are __POSTMAN__ examples



### **Feedback from MercadoLibre**
- A nivel documentación, en el entregable sólo se encontraron algunos métodos comentados. Se valoraba encontrar una introducción a la solución y una descripción de la solución tomada, como así también las suposiciones que se hicieron para llegar a la solución. __IN PROGRESS...__
- La solución en si no es correcta, aunque resuelve los niveles solicitados en el enunciado. Para el cálculo de trilateración utiliza una librería externa pero no lo hace correctamente y tampoco contempla casos de borde en el procesamiento del mensaje. Además, la solución no escala bien ante cambios en los requerimientos (numero de satelites, por ejemplo) y si bien existe validación de parámetros de entrada, no se valida, por ejemplo, que las distancias provistas terminen permitiendo formar un triángulo real.
- En cuanto a los tests, no se encontraron tests unitarios completos, solo cubren validación de parámetros y no se prueba la funcionalidad principal de los servicios. Tampoco se analizan casos de borde ni hay cálculo de cobertura.
- Teniendo todos los aspectos en cuenta, no vemos en el candidato para el seniority que estamos buscando para el equipo en este momento. De todas formas se valora el tiempo invertido en el challenge y recomendamos tener estas recomendaciones en cuenta al encarar futuros exámenes técnicos.