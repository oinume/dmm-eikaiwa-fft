import React from 'react';
import { ClipLoader } from 'react-spinners';

type Props = {
  loading: boolean;
  message: string;
  css: string;
  size: number;
};

export const Loader: React.FC<Props> = ({ loading, message, css, size }) => {
  if (size === undefined) {
    size = 40;
  }
  return loading ? (
    <div className="overlay-content">
      <div className="wrapper">
        <ClipLoader css={css} size={size} color={'#123abc'} loading={loading} />
        <span className="message">{message}</span>
      </div>
    </div>
  ) : null;
};
