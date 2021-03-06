// VARIAVEIS

// constantes não podem ser mudadas 
const idade = 21;
const nome = "Mattheus";
const sobrenome = "Russi";

// let pode ser alterada para qualquer tipo de variavel
let ano = 2021;


// ESCREVENDO NO CONSOLE
console.log(nome+ " " +sobrenome);


// MATEMATICA
console.log(10 + 4 * 3);
console.log((10 + 4) * 3);

console.log("ano " + 2020);
console.log("2" + "2");

// convertendo tipos
console.log(parseInt("2") + parseInt("2"));

// interpolando variaveis
console.log(`minha idade ${idade}`);

// ctr+k+c comenta  todo de uma vez

//LISTAS
const listadestinos = new Array(
    "Salvador",
    "São Paulo",
    "Rio de Janeiro",
    "Curitiba"
);

// adicionando elementos a lista
listadestinos.push("Florianopolis");

//removendo elementos
listadestinos.splice(3,1); //primeiro o numero do elemento depois quantos elementos
console.log(listadestinos[0]) //elemento especifico


// COMDICIONAIS
if(idade >= 18){
    console.log("voce é maior de idade");
}else{
    console.log("voce é menor de idade");
};

// shift + alt + F == formatar codigo

// OU == OR == ||
if(idade >= 18 || idade == true){
    console.log("voce é maior de idade")
}

// tambem == AND == &&
if(idade >= 18 && idade == true){
    console.log("voce é maior de idade")
}

// LOPS

// while
let contador = 0;
while(contador < 6){
    console.log(contador);
    contador ++;
};

//break
if(contador == 2){
    break
}

//FOR
for(let cont = 0 ; cont < 6; cont++){
    console.log(cont);
};

