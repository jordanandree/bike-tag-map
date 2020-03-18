import React from 'react';
import PropTypes from 'prop-types';

const Post = ({ id, title }) => {
  return (
    <div key={id} id={id}>
      {title}
    </div>
  );
};

Post.propTypes = {
  id: PropTypes.string.isRequired,
  title: PropTypes.string.isRequired,
};

export default Post;
