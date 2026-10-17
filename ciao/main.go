package main //questo dichiara che il programma fa parte del package main,
//in Go i package sono raccolte di file, nello specifico il
//package main indica il punto di inizio del programma.
//Questo package deve obbligatoriamente contenere la funzione main, altrimenti il compilatore
// va in crisi

// per avviarlo: da terminale "go run main.go"
// il comando run ha eseguito il programma dopo averlo compilato.
// GO è infatti un linguaggio compilato, con tempi di build rapidissimi
//
import "fmt" //package format, che contiene le funzioni di stampa

func main() { //quando crei una funzione, la graffa di apertura deve essere sulla stessa riga di dichiarazione
	fmt.Println("hello world!") //non è previsto ';'
	//gli elementi accessibili nei package iniziano sempre con la Maiuscola

	//2. Variabili [Types, Constant, Pointer]
	//Go è un linguaggio staticamente tipizzato,significa che il tipo delle
	//variabili è noto a tempo di compilazione e non può variare durante l'esecuzione

	//per creare una variabile si usa la keyword "var" con nome in camelcase
	var myVariable string // se la variabile definita non viene usata, il compilatore da errore
	// L questo per assicurare un codice pulito e privo di elementi inutilizzati

	//Per valorizzare una variabile, si usa ' = '
	myVariable = "gioia"
	// se voglio una stringa multi riga si usa il back tick opt + \
	myVariable = `questa
	è 
	una stringa
	multiriga`

	//per creare e valorizzare una variabile in un colpo solo:
	myVariable2 := "gioia2" // se ometto var devo mettere i : (permette inferenza nonostante sia statically types)

	fmt.Println(myVariable)
	fmt.Println(myVariable2)

	//per definire una COSTANTE invece
	const myCostant = "costante" // il type è opzionale
	fmt.Println(myCostant)

	// in Go esistono anche i puntatori che conservano l'indirizzo di memoria della variabile
	// si definiscono con &nomeDellaVariabile
	pointer := &myVariable //assegnandola ad una variabile e stampandola, vedo l'indirizzo
	fmt.Println(pointer)
	fmt.Println(*pointer) // anteponendo l'asterisco ottengo il valore della variabile puntata
	// con questa sintassi posso anche assegnargli un nuovo valore
	*pointer = "ambrogio"
	fmt.Println(*pointer)

	// 3. COLLECTIONS [ARRAY, SLICE, MAP]
	//Array: collezione ordinata di elementi dello stesso tipo
	// per dichiarare un array:
	numbers_array := [3]int{4, 3, 5} // prima il numero di elementi, poi il tipo e poi gli elementi
	fmt.Println(numbers_array)       // stampo l'intero array
	fmt.Println(numbers_array[1])    // stampo il singolo elemento accedendo all'indice

	//per modificare il valore di un elemento accedo sull'array all'indice di quell'elemento
	numbers_array[2] = 5

	// posso anche inizializzare l'array con n valori di default

	var numbers_array2 [3]int // di default per int è 0
	// e assegnargli un valore in seguito
	numbers_array2[1] = 5
	fmt.Println(numbers_array2)
	// l'array è una collezione di dimensione fissa, quindi una volta dichiarate le sue dimensioni, non posso aggiungere elementi o rimuoverli

	// se voglio una collezione dinamica uso lo slice
	numbers_slice := []int{4, 2, 6} // si dichiara come l'array ma omettendo la lunghezza
	fmt.Println(numbers_slice)
	//prevede l'aggiunta di elementi con metodo append
	numbers_slice = append(numbers_slice, 3)
	fmt.Println(numbers_slice)
	//posso ottenere il numero di elementi in una collezione con la funzione len
	fmt.Println(len(numbers_slice))
	// posso otteere la capacità di una collezione, ovvero lo spazio per essa allocato in memoria
	//chiamanso la funzione cap
	fmt.Println(len(numbers_slice), cap(numbers_slice))
	// per lo slice la capacità è maggiore della lunghezza, se continuo ad aggiungere elementi
	// la capicità verrà sempre incrementata in vista di aggiunte future
	// nell'array , essendo di dimensione fissa, la capacità è sempre uguale alla lunghezza
	fmt.Println(len(numbers_array), cap(numbers_array)) // è quindi un pò più efficente in termini di memoria

	//la MAP è infine una collezione dinamica e non ordinata di coppie chiave-valore
	// si dichiara con la kw map seguita dal tipo delle chiavi fra [] e dal tipo di valori

	age := map[string]int{
		"Alice":     23,
		"Giampirlo": 34,
	}
	fmt.Println(age["Giampirlo"]) // stampo il valore della chiave
	// se la chiave non esiste, il programma non va in errore, verrà restituito il valore di default
	fmt.Println(age["Giampirlone"])
	// posso verificare se una chiave è presente:
	val, isPresent := age["Bob"]
	fmt.Println(val, isPresent)

	// per modificare il valore di una chiave sulla map indichiamo fra [] e assegno il nuovo valore
	age["Alice"] = 28
	fmt.Println(age)
	// stessa sintassi per aggiungere una nuova coppia k-v
	age["Nino"] = 36
	fmt.Println(age)

	//per eliminare una coppia, chiamo la funzione delete
	delete(age, "Nino") // map, chiave da eliminare
	fmt.Println(age)

	//4. CYCLES [ for,for..and for!]
	// il ciclo for in go è un mix tra nuova e vecchia scuola
	// for i := 0; i > 5; i++ { -> i comandi preciclo e post iterazione sono pzionali
	// }
	i := 0
	for i > 5 {
		fmt.Println(i)
		i++
	} // questo equivale ad un ciclo while, che in go non esiste formalmente.
	//posso ottenere un loop infinito omettendo anche la condizione
	y := 0
	for {
		fmt.Println(y)
		y++
		break //comando di uscita
	}
	// in Go non esiste neanche il for-each per ciclare attraverso una collezione
	for index, valore := range numbers_array { //dichiaro 2 variabili di ciclo che conterranno indice e rispettivo valore assegnandole al range della collezione da ciclare
		fmt.Println(index, valore) // se non mi serve indice, e non lo uso, il compilatore si arrabbia
		// per ignorarlo devo mettere un _
		//for _, val := ........
		// se invece non mi serve il valore, posso ometterlo lasciando solo l'indice

	}
	//posso usare il ciclo for anche con le map, seguendo lo stesso criterio
	for key, val := range age {
		fmt.Println(key, val)
	}

}
