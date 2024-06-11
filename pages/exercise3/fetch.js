/** 
 * Exercise 4: Fetching Data
 * Erstellen Sie eine neue Datei "data.js" im Verzeichnis pages/api, die JSON-Daten zurückgibt.
   Implementieren Sie die API-Route und rufen Sie die Daten in dieser Datei ab.
 */


// Fetch data from the API route
export async function fetchData() {
  const response = await fetch('/data');
  const data = await response.json();
  return data;
}
