"""
This script runs the application using a development server.
It contains the definition of routes and views for the application.
"""

import random 
import json

from flask import Flask
app = Flask(__name__)

# Make the WSGI interface available at the top level so wfastcgi can get it.
wsgi_app = app.wsgi_app



@app.route('/')
def menu():

    #Asignación de valores a las listas
    sospechoso=["El/La mejor amigo(a)","El/la novio(a)","El/la vecino(a)","El mensajero","El extraño","El/la hermanastro(a)","El/la colega de trabajo"]
    arma=["Pistola","Cuchillo","Machete","Pala","Bate","Botella","Tubo","Cuerda"]
    motivo=["Venganza","Celos","Dinero","Accidente","Drogas","Robo"]
    cuerpo=["Cabeza","Pecho","Abdomen","Espalda","Piernas","Brazos"]
    lugar=["Sala","Comedor","Baño","Terraza","Cuarto","Garage","Patio","Balcón","Cocina"]

    #Junto todas las listas
    lista=[sospechoso,arma,motivo,cuerpo,lugar]

    #Asigno las restricciones y luego las solución junto a esas restricciones
    rest=restrictions(lista,15)
    solv=sol_validation(lista,rest)
    print("Solución:\t"+str(solv))
    return json.dumps(busqueda(lista, solv,rest))


def busqueda(lista,solv,rest):

    #Intento de solución y el número de intento
    s_try=[]
    intento=0
    registro=[]

    while s_try!=solv:

        intento+=1

        #Asigno una lista aleatoria al intento
        s_try=sol_validation(lista,rest)

        #Elimino un elemento de la lista para no volverlo a sugerir
        if s_try!=solv:

            delete=""

            while delete=="":

                #se elige el elemnto aleatorio
                i=random.randrange(0,len(s_try))
                delete=s_try[i]

                #Por si el elemento aleatorio no esta incorrecto, así, junto al while, intenta con otro hasta llegar al inocrrecto
                if delete in solv:

                    delete=""

                else:
                    #Se elimina eel elemnto incorrecto seleccionado
                    lista[i].remove(delete)
            registro=registro+[[intento,s_try,delete]]
    registro=registro+[[intento,s_try,"*"]]
    return registro


def sol_validation(lista,rest):
    aux=True
    while aux:
        aux=False
        solv=solution_select(lista)
        for x in rest:
            if x[1] in solv and x[0] in solv:
                aux=True
    return solv


def restrictions(lista, i):
    rest=[]
    for x in range(i):
        l1=lista[random.randrange(0,len(lista))]
        l2=l1
        while l1==l2:
            l2=lista[random.randrange(0,len(lista))]
        e1=l1[random.randrange(0,len(l1))]
        e2=l2[random.randrange(0,len(l2))]
        r=[e1,e2]
        rest=rest+[r]
    return rest

def solution_select(lista):
    solv=[]
    for l in lista:
        solv=solv+[l[random.randrange(0,len(l))]]
    return solv

if __name__ == '__main__':
    import os
    HOST = os.environ.get('SERVER_HOST', 'localhost')
    try:
        PORT = int(os.environ.get('SERVER_PORT', '5555'))
    except ValueError:
        PORT = 5555
    app.run(HOST, PORT)
