# Saga Orquestada en Golang

## Definición
El patrón **Saga** es un patrón de diseño para la gestión de transacciones distribuidas. Se utiliza cuando una operación compleja se divide en múltiples pasos independientes, y cada paso debe garantizar que, en caso de fallo, se ejecute una acción compensatoria para mantener la consistencia del sistema.

En una **Saga Orquestada**, un componente central (el **Orquestador**) se encarga de coordinar la ejecución de cada paso de la saga y, en caso de error, de revertir los pasos ya ejecutados mediante las funciones de compensación.

## Usos
El patrón Saga es útil en sistemas donde:
- No es posible utilizar transacciones distribuidas tradicionales.
- Se requiere garantizar la consistencia eventual de un flujo de trabajo complejo.
- Existen múltiples servicios que deben coordinarse de manera fiable.

Ejemplos comunes incluyen:
- Procesamiento de órdenes en comercio electrónico.
- Gestión de reservas en viajes (hoteles, vuelos, alquiler de autos).
- Procesos bancarios y de pago.

## Componentes
1. **Orquestador**: Controla el flujo de la saga, ejecutando los pasos en orden y manejando fallos mediante compensaciones.
2. **Pasos de la saga**: Cada paso representa una acción dentro del proceso.
3. **Compensaciones**: Funciones que revierten un paso en caso de fallo posterior.
4. **Dependencias**: Algunos pasos pueden depender de la ejecución exitosa de otros.

## Aplicación en el Mundo Real
El uso de sagas orquestadas es común en arquitecturas de microservicios donde cada servicio maneja una parte del proceso. Por ejemplo, en un sistema de pedidos en línea:
1. Reservar el inventario.
2. Cobrar al cliente.
3. Confirmar el pedido.

Si el cobro falla, el sistema debe liberar el inventario. Aquí es donde una saga orquestada es clave para mantener la integridad del proceso.

## Implementación en Golang
Este repositorio contiene un ejemplo de implementación de una saga orquestada en Golang, donde se manejan dependencias entre pasos y compensaciones en caso de fallos.

### Ejecución
Para ejecutar el código, simplemente compílalo y ejecútalo con:
```sh
go run main.go
```

El código simula un flujo de trabajo con tres pasos, en el que el segundo paso falla intencionalmente, activando el mecanismo de rollback.

## Conclusión
El patrón Saga Orquestada es una solución eficaz para la gestión de transacciones distribuidas sin la necesidad de bases de datos transaccionales tradicionales. Su implementación en Golang permite construir sistemas resilientes y confiables.
