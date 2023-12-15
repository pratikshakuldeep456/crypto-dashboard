document.addEventListener('DOMContentLoaded', function() {
    fetchWazirXData();
});

function fetchWazirXData() {
    fetch('http://localhost:8080/api/wazirx')
        .then(response => response.json())
        .then(data => {
            let output = '<h2>WazirX Tickers</h2>';
            data.forEach(ticker => {
                output += `
                    <p><strong>${ticker.baseAsset}/${ticker.quoteAsset}</strong>: ${ticker.lastPrice}</p>
                `;
            });
            document.getElementById('data').innerHTML = output;
        })
        .catch(error => console.error('Error:', error));
}
