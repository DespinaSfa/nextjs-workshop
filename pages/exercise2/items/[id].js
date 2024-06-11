import { items } from '../../data';
import styles from '../../../styles/Item.module.css';

export async function getStaticPaths() {

  return {
    paths,
    fallback: false,
  };
}

export async function getStaticProps({ params }) {
  const item = items.find((item) => item.id.toString() === params.id);
  return {
    props: {
    },
  };
}

export default function ItemPage({ item }) {
  return (
    <div className={styles.container}>
    <h1 className={styles.title}>{item.name}</h1>
    <p className={styles.description}>{item.description}</p>
  </div>
  );
}
