// CLASE BASE
class Cliente{
    nome;
    cpf;
    agencia;
}

const cliente1 = new Cliente();

cliente1.nome = 'pedro';
cliente1.cpf = 08812267689;
cliente1.agencia = 1001;

console.log(cliente1);

// definindo funções
class ContaCorente{
    agencia;
    saldo;

    sacar(valor){
        if (this.saldo >= valor){
            this.saldo -= valor;
        }
    }
}
