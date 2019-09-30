// a comment

class Person {
    constructor() {
        this.a = 'a';
        this.b = 3;
    }
    updateA(a) {
        this.a = a
    }
}

let cody = new Person;

const someValues = ['a', 'b', 'c'];

const testFunction = () => {
    if (3>2) {
        console.log('this is true')
    } else {
        console.log('impossible')
    }
}

document.getElementById('somehtmlid').addEventListener("click", function(e){ 
    console.log('this is an event listener')
});
