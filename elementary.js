function multiply (a,b) {
    let coun = 0;
    for (let i = 1; i <= b; i++) {
        coun += a
    };
    return coun;
};
function divide(a,b) {
    let count = 0;
    while (a >= b) {
        a -= b;
        count++;
    }
    return count;
}
function modulo(a,b) {
    while (a >= b) {
        a -= b;
    }
    return a
}
console.log(modulo(22,4));
