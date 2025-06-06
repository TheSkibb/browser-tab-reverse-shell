const url = "http://localhost:8080/get"; const r1 = await fetch(url);const r2 = await r1.text(); eval(r2);

