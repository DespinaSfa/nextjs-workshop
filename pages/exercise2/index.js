/** 
 * Exercise 2:
 * Vervollständigen Sie den Code hier und in [id].js, um eine Artikelliste und eine Artikeldetailseite 
 * in einer Next.js-Anwendung zu erstellen.
 */

import Link from 'next/link';
import styles from '../../styles/Home.module.css';
import { items } from '../data';

export async function getStaticProps() {
  return {
    props: {
    },
  };
}

export default function HomePage({ items }) {
  return (
    <div className={styles.container}>
      <h1 className={styles.title}>Artikelliste</h1>
      <ul className={styles.list}>
        {items.map((item) => (
          <li key={item.id} className={styles.listItem}>
            <Link href={`exercise2/items/${item.id}`}>
              {item.name}
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
}
