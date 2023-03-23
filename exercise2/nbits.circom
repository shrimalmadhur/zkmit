pragma circom 2.1.4;

include "circomlib/poseidon.circom";
// include "https://github.com/0xPARC/circom-secp256k1/blob/master/circuits/bigint.circom";

template Num2Bits(nBits) {
    signal input in;
    signal b[nBits];

    // assign 
    for (var i = 0; i < nBits; i++) {
        b[i] <-- (in \ (2 ** i)) % 2;
    }
    
    var sum = 0;

    // calculate final results
    for (var i; i < nBits; i++) {
        sum += b[i] * (2 ** i);
    }

    sum === in;

    // asset bit values
    for (var i = 0; i < nBits; i++) {
        0 === b[i] * (b[i] - 1);
    } 
}

component main = Num2Bits(5);

/* INPUT = {
    "in": "11"
} */