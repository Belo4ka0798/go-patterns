## Строитель (Builder)

Паттерн Строитель - порождающий шаблон проектирования, предоставляет способ создания составного объекта.

Паттерн строитель отделяет конструирование сложного объекта от его представления так, что в результате одного и того же процесса конструирования могут получаться разные представления.

![image](Builder-1.gif)

Для реализации этого паттерна на описаном примере требуется:

1. Класс Director, который будет распоряжаться строителем и отдавать ему команды в нужном порядке, а строитель будет их выполнять;
2. Базовый абстрактный класс Builder, который описывает интерфейс строителя, те команды, которые он обязан выполнять;
3. Класс ConcreteBuilder, который реализует интерфейс строителя и взаимодействует со сложным объектом;
4. Класс сложного объекта Product.

Подведем итог.

Паттерн Builder определяет процесс поэтапного построения сложного продукта. После того как будет построена последняя его часть, продукт можно использовать.