import React from 'react';

import Message from './Message';

const Messages = ({ msgs }) => {
  return (
    <div className="messages">
      {msgs.map((msg, index) => <Message msg={msg} key={index} />)}
    </div>
  );
};

export default Messages;
