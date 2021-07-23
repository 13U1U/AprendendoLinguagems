import re

endereco = "rua antoni luiz moura gonzaga, porto da lagoa, 23440-120 apocalipise now"


cpe = re.compile("[0-9]{5}[-]?[0-9]{3}")

# retorna objeto Metch
busca = cpe.search(endereco) # caso não encontre retorna o valor NONE

if busca:
    resultado = busca.group()
    print(resultado)


# [0-9] se achar qualquer coisa que está no colchete
# (abc) tem que satisfazer esatamente o que tá em chaves
# ? pode ou não conter a informação 
# {4} quantidade de vezes que aquela informação tem que aparecer
