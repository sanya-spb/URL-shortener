# URL-shortener
Backend-development on Go. Module 1

## ТЗ
### Введение
Данный курс является вводным в разработку веб-приложений на Go, основная цель этого курса - дать общее представление о процессах разработки и поддержки современных веб-сервисов, научить работе с актуальными инструментами и практиками.  

В процессе выполнения домашних заданий мы поэтапно будем собирать итоговый курсовой проект - сервис-сокращатель ссылок (URL shortener), на примере которого мы и познакомимся в деталях с тем, как устроены современные веб-приложения.  

Разобравшись с ключевыми компонентами современных веб-приложений, мы можем перейти к поэтапному планированию нашего будущего курсового проекта. Для начала, нам необходимо правильно сформулировать стоящую перед нами задачу, а для этого, разберемся, что такое URL shortener.  

**URL shortener** (“сокращатель ссылок”) - это сервис, позволяющий пользователю генерировать  для произвольного URL’a его короткую версию, которую удобно вставлять в различные публикации, сообщения, новости, промо-материалы и так далее. Также сервис позволяет получать статистику переходов по каждому сгенерированному URL’у, что будет полезно, если его владелец захочет узнать сколько людей перешло по короткой ссылке из конкретного источника.  

Возможно, вы знакомы с сервисам типа https://bitly.com/. Это и есть сокращатели ссылок.  

*Как может работать такой сервис?*  
Переходя по короткой ссылке, пользователь отправляет запрос на backend, который, в свою очередь, ищет в своем хранилище (например, в базе данных) длинную версию того короткого URL, по которому был совершен переход. В случае нахождения в БД искомой записи, сервер перенаправляет пользователя на соответствующий длинный URL и фиксирует в базе факт перехода для того, чтобы в дальнейшем предоставить статистику владельцу URL’а.  

*Как создаются короткие URL и как по ним получить статистику?*  
Например, веб-приложение помимо серверной части, которая будет непосредственно обрабатывать переходы по коротким ссылкам, также может иметь и frontend, где любой желающий сможет генерировать короткий URL с помощью html-формы и условной кнопки “Generate!”, формирующей запрос к серверу, в ответ получая  сгенерированную короткую версию и специальную ссылку, по которому он сможет получить статистику для этого URL.  

Разобьем веб-приложение на компоненты согласно сформулированным требованиям:  
* **Frontend** - пользовательский интерфейс будет осуществлять формирование запросов для генерации коротких ссылок, а также выводить статистику. 
* **Backend** - будет генерировать уникальные короткие URL, обрабатывать переходы по коротким ссылкам, создавать необходимые записи в базе для обоих случаев, а также выдавать статистику переходов.
* **Хранилище** - здесь мы будем хранить непосредственно короткие URL и соответствующие им длинные версии, а также статистику переходов: URL, по которому был совершен переход, IP пришедшего пользователя, дата и время перехода.  

В рамках данного курса мы будем заниматься реализацией на Go серверной части URL shortener’a. Взаимодействие с фронтендом будет происходить через заранее описанный REST API, с хранилищем - через интерфейс, заданный внутри приложения. 

### Рекомендации
1. Для удобства рекомендуем для разработки приложения завести отдельный репозиторий. Для сдачи курсового нужно будет предоставить ссылку на этот репозиторий, включающий в себя всю реализованную функциональность.
1. По ходу написания кода приложения постарайтесь сразу же писать тесты на реализованную логику.
1. Сдача курсового предполагается в конце курса, но желательно начать писать его заранее. В рамках курса у нас будет несколько уроков-консультаций, на которых мы разберем вопросы, возникающие по ходу выполнения проекта. Чтобы такие консультации проходили максимально эффективно, как только у вас будут появляться промежуточные результаты на проверку, обязательно делитесь ими с преподавателем.
1. В рамках курсового проекта оцениваться будет реализация backend-составляющей. Наличие frontend’а и базы данных не обязательно. По желанию сейчас или позднее вы также можете добавить frontend и работу с базой данных. Такой вклад поможет сделать ваше портфолио еще лучше. 

### Ход выполнения проекта
1. Подготовить краткое описание (README) проекта. Добавить туда информацию о назначении приложения и принципы его работы.
1. Продумать структуру приложения. Из каких пакетов оно будет состоять?
    * В подготовке структуры поможет материал урока “Принципы структурирования Go-приложений” из курса “Лучшие практики разработки Go-приложений”
1. Продумать, как приложение будет работать с хранилищем данных. Мы пока не работаем с базами данных, поэтому в качестве простейшего хранилища могут выступать, например, файлы.
    * Подумайте, какие действия приложение будет производить с хранилищем. В какие моменты эти действия будут производиться? Можете ли вы выделить общий интерфейс, независимый от типа хранилища?
1. Продумать HTTP и REST API интерфейс. Из каких эндпоинтов будет состоять приложение? Задокументировать принятое решение в файле README.
1. Подготовить репозиторий в соответствии с ранее изученными лучшими практиками. Внедрить автоматическую проверку кода линтерами и автоматический запуск тестов.
    * В подготовке этого пункта помогут материалы уроков предыдущих курсов - “Go. Уровень 2” и “Лучшие практики разработки Go-приложений”
1. Задокументировать и реализовать бизнес-логику сервиса по сокращению ссылок. Не забываем писать тесты!
1. Написать Dockerfile для запуска сервиса в Docker.
    * Для выполнения задания обратитесь к материалам урока “Особенности докеризации Go-приложений” текущего курса.
1. Развернуть сервис на Heroku или любой другой платформе. Внедрить практики CI/CD.
    * Для выполнения задания обратитесь к материалам урока “Сборка и развёртывание приложения. CI/CD” текущего курса.

### Критерии приема курсового проекта
1. Необходимо предоставить ссылки на репозиторий с кодом приложения (например, на GitHub) и развернутый экземпляр приложения (например, в Heroku).
1. Приложение должно реализовывать бизнес-логику в соответствии с заданием. 
1. Приложение должно быть документировано.
1. Код приложения должен быть покрыт тестами.
1. Приложение должно соответствовать лучшим практикам, изученным ранее. Проверьте, что вы учли следующие практики и принципы:
    * Запуск и завершение приложение. Чтение конфигурации и Graceful Shutdown.
    * Логирование. Приложение должно писать логи.
    * Структурирование. Структура приложения должна быть простой и понятной.
    * Наличие автоматических проверок кода с использованием линтеров.
    * Наличие Makefile’а или любого другого инструмента сборки и запуска.
1. В репозитории должен быть Dockerfile для контейнеризации приложения.
1. Для репозитория должны быть заданы Continuous Integration и Continuous Deployment для непрерывной интеграции и развертывания приложения.

## Решение


REST API: https://app.swaggerhub.com/apis/sanya-spb/URL-shortener/1.0.0#/developers/getStat

## Вопросы на Урок 10 (Консультация по курсовому проекту)
