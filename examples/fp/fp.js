function addTwo(x) {
  return x + 2;
}

function minusOne(x) {
  return x - 1;
}

function calculate(x, fn) {
  return fn(x);
}

let ans = 50;

ans = calculate(ans, addTwo);
ans = calculate(ans, minusOne);

console.log(ans);
