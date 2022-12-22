import React, { useState } from 'react';

const InputMessage = ({ content, setContent, receiver, setReceiver, handleSendMessage }) => {
  return (
    <div className="input">
      <input type="text" placeholder="Type message" value={content} onChange={(e) => setContent(e.target.value)} className="contentMsg" />
      <input type="text" placeholder="Type receiver" value={receiver} onChange={(e) => setReceiver(e.target.value)} className="receiverMsg" />
      <button onClick={handleSendMessage}>Send</button>
    </div>
  );
};

export default InputMessage;
