import React, { useEffect, useRef } from 'react';

const Message = ({ msg }) => {
  const ref = useRef()

  useEffect(() => {
    ref.current?.scrollIntoView({
        behavior: "smooth"
    })
  }, [msg])

  return (
    <div className={msg.owner ? 'message owner' : 'message'}  ref={ref}>
      <div className="messageContent">
        {
          msg.owner ?
            (<div style={{ display: 'flex', alignItems: 'center' }}>
              <p>{'Owner message | To ' + msg.receiver}<br /><br />{msg.content}</p>
            </div>
            ) :
            (msg.type == "DIRECT" ?
              <p>{'Direct message | From ' + msg.sender}<br /><br />{msg.content}</p>
              :
              <p>{'Group message | From ' + msg.groupname + ' | ' + msg.sender}<br /><br />{msg.content}</p>
            )
        }
      </div>
    </div>
  );
};

export default Message;
