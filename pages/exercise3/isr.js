/** 
 * Exercise 3: ISR
 * Verwenden Sie ISR, um Inhalte basierend auf einem Zeitintervall oder anderen Auslösern zu aktualisieren.
   Stellen Sie sicher, dass die Seite statisch generiert wird und nach einer bestimmten Zeitspanne automatisch aktualisiert wird.
 */


const ISRPage = ({ posts }) => {
  return (
    <div>
      <h1>Posts</h1>
      <ul>
        {/* Add code here to render posts. Hint: Title and body of each post should be displayed */}
      </ul>
    </div>
  );
};

export async function getStaticProps() {
  const url = 'https://jsonplaceholder.typicode.com/posts';
  {/* Implement getStaticProps */}

  return {
    props: {
    },
  };
}

export default ISRPage;
