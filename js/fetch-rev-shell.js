const url = "http://localhost:8080/"; const r1 = await fetch(url + "get");const r2 = await r1.text(); eval(r2);

