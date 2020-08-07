import React from 'react';
import styles from './index.less';

export interface FilterItemProps {
  label: string;
}
const FilterItem: React.SFC<FilterItemProps> = ({ label = '', children }) => {
  const labelArray = label.split('');
  return (
    <div className={styles.filterItem}>
      {labelArray.length > 0 && (
        <div className={styles.labelWrap}>
          {labelArray.map((item, index) => (
            <span className="labelText" key={index}>
              {item}
            </span>
          ))}
        </div>
      )}
      <div className={styles.item}>{children}</div>
    </div>
  );
};

export default FilterItem;
