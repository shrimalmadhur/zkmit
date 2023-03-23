pragma circom 2.1.4;

include "circomlib/poseidon.circom";
// include "https://github.com/0xPARC/circom-secp256k1/blob/master/circuits/bigint.circom";

template IsZero() {
    signal input in;
    signal output out;
    signal inv;

    inv <-- in != 0 ? 1/in : 0;
    out <== -inv*in + 1;
    in*out === 0;
}

component main = IsZero();

/* INPUT = {
    "in": "12"
} */