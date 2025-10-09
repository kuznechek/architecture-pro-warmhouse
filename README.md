# Задание 1. Анализ и планирование

### 1. Описание функциональности монолитного приложения

**Управление отоплением:**

- Пользователи могут удалённо включать/выключать отопление в своих домах.
- Система поддерживает добавление/удаление/обновление данных о датчиках температуры

**Мониторинг температуры:**

- Пользователи могут просматривать текущую температуру в своих домах через веб-интерфейс.
- Система получает данные о температуре с датчиков, установленных в домах.

### 2. Анализ архитектуры монолитного приложения

- Язык программирования: Go
- База данных: PostgreSQL
- Архитектура: монолитная, все компоненты системы (обработка запросов, бизнес-логика, работа с данными) находятся в рамках одного приложения.
- Взаимодействие: синхронное, запросы обрабатываются последовательно.
- Масштабируемость: ограничена, так как монолит сложно масштабировать по частям.
- Развертывание: требует остановки всего приложения.

### 3. Определение доменов и границы контекстов

**Домен:** управление датчиками (включение, выключение, изменение локации)

**Контекст работы с устройствами:**
1. сущности: датчик температуры
2. объекты-значения: управляющая команда
3. агрегаты: список датчиков
4. репозитории: репозиторий датчиков
5. сервисы: сервис управления температурой (temperature_service.go)

**Домен:** получение данных с датчиков

**Контекст работы с устройствами:**
1. сущности: датчик температуры (или другой датчик)
2. объекты-значения: дто со значением датчика, статус датчика (активен, неактивен)
3. агрегаты: список датчиков
4. репозитории: репозиторий датчиков
5. сервисы: сервис телеметрии

### **4. Проблемы монолитного решения**
-Любые изменения в текущем ПО увеличивают риск возникновения ошибок, так как все компоненты программы сильно связаны друг с другом.
**Как итог:** масштабирование и доработка повлекут непредвиденное ранее поведение программы, так что потребуется дополнительное тестирование

- Текущий продукт невозможно протестировать по частям
**Как итог:** введение каждой новой функции потребует нового тестирования всего цикла развёртывания

- Текущий продукт трудно масштабировать из-за сильной связности: для добавления нового функционала придётся переписывать имеющийся сервис или разделять его, кроме того основной цикл работы сервера также будет подвержен изменениям.
**Как итог:** масштабирование повлечёт изменение во всех областях ПО

- Неоптимизированная работа команды из-за всех вышеперечисленных факторов
**Как итог:** увеличивается время на разработку и тестирование

### 5. Визуализация контекста системы — диаграмма С4

![Warm House Context Diagram](https://github.com/kuznechek/architecture-pro-warmhouse/blob/warmhouse/schemas/context/context.png)

[Warm House Context Diagram](https://github.com/kuznechek/architecture-pro-warmhouse/blob/warmhouse/schemas/context/context.puml)

# Задание 2. Проектирование микросервисной архитектуры

**Диаграмма контейнеров (Containers)**

![Warm House System Container Diagram](https://github.com/kuznechek/architecture-pro-warmhouse/blob/warmhouse/schemas/containers/containers.png)

[Warm House System Container Diagram](https://github.com/kuznechek/architecture-pro-warmhouse/blob/warmhouse/schemas/containers/containers.puml)

**Диаграмма компонентов (Components)**

![Auth/Registration Component Diagram](https://github.com/kuznechek/architecture-pro-warmhouse/blob/warmhouse/schemas/components/container_AuthReg.png)

[Auth/Registration Component Diagram](https://github.com/kuznechek/architecture-pro-warmhouse/blob/warmhouse/schemas/components/container_AuthReg.puml)

![Device Component Diagram](https://github.com/kuznechek/architecture-pro-warmhouse/blob/warmhouse/schemas/components/container_Device.png)

[Device Component Diagram](https://github.com/kuznechek/architecture-pro-warmhouse/blob/warmhouse/schemas/components/container_Device.puml)

**Диаграмма кода (Code)**

![Device Control Code Diagram](https://github.com/kuznechek/architecture-pro-warmhouse/blob/warmhouse/schemas/code/device_control_component.png)

[Device Control Code Diagram](https://github.com/kuznechek/architecture-pro-warmhouse/blob/warmhouse/schemas/code/device_control_component.puml)

![Sensors Telemetry Code Diagram](https://github.com/kuznechek/architecture-pro-warmhouse/blob/warmhouse/schemas/code/sensors-telemetry-component.png)

[Sensors Telemetry Code Diagram](https://github.com/kuznechek/architecture-pro-warmhouse/blob/warmhouse/schemas/code/sensors-telemetry-component.puml)

# Задание 3. Разработка ER-диаграммы

![ER Diagram](https://github.com/kuznechek/architecture-pro-warmhouse/blob/warmhouse/schemas/er/er.png)

[ER Diagram](https://github.com/kuznechek/architecture-pro-warmhouse/blob/warmhouse/schemas/er/er.puml)

# Задание 4. Создание и документирование API

### 1. Тип API

- В данном случае взаимодействие между фронтеном и бекендом будет реализовано через REST API, так как у приложения простые коммуникационные потребности.

### 2. Документация API

[Device Control swagger.json](https://github.com/kuznechek/architecture-pro-warmhouse/tree/warmhouse/apps/microservice-architecture/device-control-service/swagger.json)

[Telemetry swagger.json](https://github.com/kuznechek/architecture-pro-warmhouse/tree/warmhouse/apps/microservice-architecture/sensors-telemetry-service/swagger.json)

# Задание 5. Работа с docker и docker-compose

- Create Sensor
- Get All Sensors

При каждом вызове отображается разное значение температуры.

# **Задание 6. Разработка MVP**

Необходимо создать новые микросервисы и обеспечить их интеграции с существующим монолитом для плавного перехода к микросервисной архитектуре. 

В рамках задания были созданы два микросервиса - для управления устройствами (**device-control-service**, Python) и для отслеживания приходящих запросов с устройств, оповещающих об изменении состояния устройства (**sensors-telemetry-service**, Go). Сервисы работают параллельно и не зависят друг от друга.

Для тестирования их работы можно воспользоваться Postman коллекциями:

[DeviceControl.postman_collection.json](https://github.com/kuznechek/architecture-pro-warmhouse/tree/warmhouse/apps/microservice-architecture/postman-collections/DeviceControl.postman_collection.json)

[Telemetry.postman_collection.json](https://github.com/kuznechek/architecture-pro-warmhouse/tree/warmhouse/apps/microservice-architecture/postman-collections/Telemetry.postman_collection.json)

Каждый сервис может быть масштабирован и доработан до более приближённой модели к реальным условиям (например, в данной версии нет эмуляторов устройств, и данные существуют только в формате таблицы в базе данных).
