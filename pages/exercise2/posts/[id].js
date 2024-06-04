/** 
 * Exercise 2: Datenabruf
 * Implementieren Sie getStaticPaths und getStaticProps, um Daten für die dynamischen Routen abzurufen und zu rendern.
   Stellen Sie sicher, dass die Daten zur Build-Zeit abgerufen werden und zur statischen Generierung der Seiten verwendet werden.
 */



import { useRouter } from 'next/router';

const Post = ({ post }) => {
  const router = useRouter();

  if (router.isFallback) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h1>Title</h1>
      <p>Body</p>
    </div>
  );
};

export async function getStaticPaths() {
  const response = await fetch('https://jsonplaceholder.typicode.com/posts');
  const posts = await response.json();

  /** Insert code 
   * Use String(...id) to convert the id to a string 
   * 
   */

  return {
    paths,
    fallback: true 
  };
}

export async function getStaticProps({ params }) {
  const response = await fetch(`https://jsonplaceholder.typicode.com/posts/${params.id}`);
  const post = await response.json();

  return {
    props: {
      post
    },
    revalidate: 60
  };
}

export default Post;
