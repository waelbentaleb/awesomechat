import React, { useState } from 'react';
import '../style.scss';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faCircle } from '@fortawesome/free-solid-svg-icons'

const Sidebar = ({
  username, handleCreateUser, setUsername,
  createGroupName, setCreateGroupName, handleCreateGroup,
  joinGroupName, setJoinGroupName, handleJoinGroup,
  leftGroupName, setLeftGroupName, handleLeftGroup,
  channels, connectedUsername
}) => {

  return (
    <div className="sidebar">
      <div className="connectButtons">
        {
          connectedUsername ? <h4 className='connectedUsers'> Welcome {connectedUsername} 🎉</h4> :
            <h4 className='connectedUsers'> You need to connecet</h4>
        }
        <br />

        <div className="input">
          <input type="text" placeholder="Enter your username" value={username} onChange={(e) => setUsername(e.target.value)} />
          <button onClick={handleCreateUser}>Connect</button>
        </div>

        <br />
        <br />

        <h4 className='connectedUsers'>Groups section</h4>

        <div className="input" style={{ marginTop: "1.5em" }}>
          <input type="text" placeholder="Enter group name" value={createGroupName} onChange={(e) => setCreateGroupName(e.target.value)} />
          <button onClick={handleCreateGroup}>Create</button>
        </div>

        <div className="input">
          <input type="text" placeholder="Enter group name" value={joinGroupName} onChange={(e) => setJoinGroupName(e.target.value)} />
          <button onClick={handleJoinGroup}>Join</button>
        </div>

        <div className="input">
          <input type="text" placeholder="Enter group name" value={leftGroupName} onChange={(e) => setLeftGroupName(e.target.value)} />
          <button onClick={handleLeftGroup}>Leave</button>
        </div>
      </div>

      <br />
      <br />

      <div>
        <h4 className='connectedUsers'>Connected channels</h4>
        <div className="channels">
          {channels.map((item, index) => <p key={index}><FontAwesomeIcon icon={faCircle} className="connectedCircle" />  {item.type} - {item.identifier} </p>)}
        </div>
      </div>
    </div>
  );
};

export default Sidebar;
