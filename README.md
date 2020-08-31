# GENOVA CHALLENGE

**Start Locally**

To exec app run "go run ." on working dir of the project


**Use Public URL (Google App Engine)**
- TopSecret service --> http://genovachallenge.uc.r.appspot.com:80/topsecret
- TopSecretSplit service --> http://genovachallenge.uc.r.appspot.com:80/topsecret_split

Use some of the examples inside the folder "examples" to test the app. They are __POSTMAN__ examples



### **Feedback from MercadoLibre**
- A nivel documentación, en el entregable sólo se encontraron algunos métodos comentados. Se valoraba encontrar una introducción a la solución y una descripción de la solución tomada, como así también las suposiciones que se hicieron para llegar a la solución.
- La solución en si no es correcta, aunque resuelve los niveles solicitados en el enunciado. Para el cálculo de trilateración utiliza una librería externa pero no lo hace correctamente y tampoco contempla casos de borde en el procesamiento del mensaje. Además, la solución no escala bien ante cambios en los requerimientos (numero de satelites, por ejemplo) y si bien existe validación de parámetros de entrada, no se valida, por ejemplo, que las distancias provistas terminen permitiendo formar un triángulo real.
- En cuanto a los tests, no se encontraron tests unitarios completos, solo cubren validación de parámetros y no se prueba la funcionalidad principal de los servicios. Tampoco se analizan casos de borde ni hay cálculo de cobertura.
- Teniendo todos los aspectos en cuenta, no vemos en el candidato para el seniority que estamos buscando para el equipo en este momento. De todas formas se valora el tiempo invertido en el challenge y recomendamos tener estas recomendaciones en cuenta al encarar futuros exámenes técnicos.