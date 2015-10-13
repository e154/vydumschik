Vydumschik (Выдумщик)
---------------------

`Vydumschik`- пакет, выдумывающий случайные данные для наполнения баз данных. Вдохновлен гемом [`vydumschik`](https://github.com/leonid-shevtsov/vydumschik).

### Установка

```bash
go get github.com/e154/vydumschik
```

### Имена 

```bash
vydumschik.First_name()
Иларион

vydumschik.Middle_name()
Алексеевич
	
vydumschik.Sur_name()
Ильин

vydumschik.Full_name()
Кудряшова Анжела Мефодиевна
```

### Адреса

```bash
address := new(vydumschik.Address)

address.Street()
ул. Физкультурная

address.Street_address()
ул. Детская, д. 38, кв. 40

address.Street_address()
ул. Краснодонская, д. 41, кв. 82

address.Street_address()
ул. Гражданская, 3/75
```

### Lorem ipsum

```bash
lorem := new(vydumschik.Lorem)
lorem.Worlds(5)
Integer lacinia lacinia erat Nulla. 

for i:=0;i<3;i++ {
lorem.Speech(1)
}
Ut erat ipsum, faucibus sit amet, interdum in, tempor rutrum, metus.
Donec viverra pede ac tortor.
Nunc felis.

lorem.Paragraphs(2)
Cras quis libero ut urna pulvinar mattis. Vestibulum suscipit gravida nulla. Proin condimentum sapien ut augue. Cras turpis ligula, pharetra in, malesuada at, eleifend vel, pede. Phasellus turpis nisi, placerat in, semper sed, tincidunt a, tortor. Vivamus est nibh, mattis vitae, aliquet non, dignissim at, velit. Etiam elementum odio non erat. Donec varius felis sed nisl. Vestibulum imperdiet. Donec viverra pede ac tortor.

Maecenas ullamcorper leo et libero. Proin nec eros. In id lorem eu tellus tempus bibendum. Curabitur eu odio. Curabitur fermentum sem ut arcu. Ut eget leo ultricies mauris malesuada sollicitudin. Vivamus fringilla auctor nisl. Curabitur venenatis. Duis tempus, nunc quis convallis posuere, eros nibh lacinia nisl, non molestie enim lorem at nisi. Aenean dui. Cras elementum felis sit amet tellus consequat vehicula. Mauris rhoncus. Donec posuere mauris adipiscing tortor. Curabitur sagittis leo in magna.


str := lorem.Bytes(112)
Donec varius felis sed nis. Morbi sed metus quis odio luctus lacini. Proin condimentum sapien ut augu. Nunc feli

len(str)
112
```

### LICENSE

Vydumschik is licensed under the [MIT License (MIT)](https://github.com/e154/vydumschik/blob/master/LICENSE).