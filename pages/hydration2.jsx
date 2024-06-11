
import { useState, useEffect } from 'react';
import styles from '../styles/Hydration.module.css';

export default function HydrationExample() {
  const [count, setCount] = useState(0);
  const [hydrated, setHydrated] = useState(false);

  useEffect(() => {
    const timer = setTimeout(() => {
      setHydrated(true);
    }, 1000); //  delay 1 second

    return () => clearTimeout(timer); 
  }, []); 

  return (
    <div className={styles.container}>
      <h1>Hydration Beispiel</h1>
      <button
        className={styles.button}
        onClick={() => setCount(count + 1)}
        style={{ visibility: hydrated ? 'visible' : 'hidden' }}
      >
        Click me! Count: {count}
      </button>
    </div>
  );
}
