MetricStructures
=====================
Этот репозиторий представляет из себя совокупность структур с метриками для хаба и майнера.
Структуры состоят из метрик трех типов: структурные, финансовые, технические.
Эти структуры представляют из себя основу для создания системы фильтров, которые являются 
основой рейтинговой системы.
Результаты рейтинговой системы используются в наборе данных для алгоритмов машинного обучения, которые
в последствии смогут выполнять функцию прогнозирования для хаба и майнера, упрощая выбор и увеличивая
уровень доверия к системе и рейтингам.

## Структурные метрики

### Структурные метрики для хаба
1) HubAddress -  фиксируется адрес хаба.
2) HubPing - использование ping в качестве диагностики сети в запросах скорости соединения хаба. Определяет факт потери пакетов отраженного сигнала от источника соединения.
3) HubService - количество доступных для хаба сервисов (этого пока не работает).
4) CreationDate - дата активации (регистрации) хаба. Определяет как давно хаб зарегестрирован (для оценки его активности в рамках определенного промежутка времени).
5) HubLifeTime - время жизни хаба.
6) PayDay - атрибут определяет количество средств, которое может предложить хаб.
7) FreezeTime - общее количество времени, на которое хаб был заморожен.
8) AmountFreezeTime - количество фактов заморозки хаба.
9) TransferLimit - лимит на отправку (_value).
10) SuspectStatus - статус становится true, если хаб становится подозреваемым в мошенничестве.
11) DayLimit - лимит средств, которые хаб может отправить.
12) AvailabilityPresale - атрибут определения наличия пресейл токенов у хаба.
13) SpeedConfirm -  атрибут определят скорость отклика хаба, влияет на увеличение вероятности активности хаба.
14) HubStack - атрибут определяет сколько участники холдят на своих кошельках.

### Структурные метрики для майнера
1) MinAddress -  фиксируется адрес майнера.
2) MinPing - спользование ping в качестве диагностики сети в запросах скорости соединения майнера. Определяет факт потери пакетов отраженного сигнала от источника соединения.
3) MinStack - атрибут определяет сколько участники холдят на своих кошельках.
4) CreationDate - дата активации (регистрации) майнера. Определяет как давно майнер зарегестрирован (для оценки его активности в рамках определенного промежутка времени).
5) MinService - количество доступных для майнера сервисов (этого пока не работает).

## Технические метрики
 
## Финансовые метрики