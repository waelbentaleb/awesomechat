import React, { useState, useEffect, useCallback } from 'react';
import { ToastContainer, toast } from 'react-toast'
import './style.scss';

import Sidebar from './components/Sidebar';
import Messages from './components/Messages';
import InputMessage from './components/InputMessage';

import { createUser, connectStream, sendMessage, createGroupChat, joinGroupChat, leftGroupChat, listChannels } from './services/chatClient'

export default function App() {
  const [username, setUsername] = useState('')
  const [chatStream, setChatStream] = useState(null)

  const [connectedUsername, setConnectedUsername] = useState('')

  const [content, setContent] = useState('')
  const [receiver, setReceiver] = useState('')
  const [msgs, setMsgs] = useState([])

  const [createGroupName, setCreateGroupName] = useState('')
  const [joinGroupName, setJoinGroupName] = useState('')
  const [leftGroupName, setLeftGroupName] = useState('')

  const [channels, setChannels] = useState([])


  const handleCreateUser = async () => {
    try {
      await createUser(username)
      const returnedStream = await connectStream(username)
      setChatStream(returnedStream)

      toast.success("Welcome to Awesome chat ^_^")
      setUsername('')
      setConnectedUsername(username)
    } catch (err) {
      toast.error(err.message)
    }

  }

  const handleSendMessage = async () => {
    try {
      await sendMessage(content, receiver)
      const msg = {
        owner: true,
        content: content,
        receiver: receiver,
      }

      setMsgs([...msgs, msg])
      setContent('')
      setReceiver('')
    } catch (err) {
      toast.error(err.message)
    }
  }

  const handleCreateGroup = async () => {
    try {
      await createGroupChat(createGroupName)
      toast.success("Group created successfully")
      setCreateGroupName('')
    } catch (err) {
      toast.error(err.message)
    }
  }

  const handleJoinGroup = async () => {
    try {
      await joinGroupChat(joinGroupName)
      toast.success("Group joined successfully")
      setJoinGroupName('')
    } catch (err) {
      toast.error(err.message)
    }
  }

  const handleLeftGroup = async () => {
    try {
      await leftGroupChat(leftGroupName)
      toast.success("Group lefted successfully")
      setLeftGroupName('')
    } catch (err) {
      toast.error(err.message)
    }
  }

  const updateState = useCallback(async () => {
    const response = await listChannels()
    const list = response.itemsList.sort(customCompare)
    setChannels(list)
  }, []);

  useEffect(() => {
    setInterval(updateState, 3000);
  }, [updateState]);

  if (chatStream) {
    chatStream.on("data", (response) => {
      const newMsg = response.toObject()
      setMsgs([...msgs, newMsg])
    });

    chatStream.on("status", function (status) {
      console.log(status.code, status.details, status.metadata);
    });

    chatStream.on("end", () => {
      console.log("Stream ended.");
    });
  }

  return (
    <div>
      <div className="home">
        <div className="container">
          <Sidebar
            username={username}
            setUsername={setUsername}
            handleCreateUser={handleCreateUser}

            createGroupName={createGroupName}
            setCreateGroupName={setCreateGroupName}
            handleCreateGroup={handleCreateGroup}

            joinGroupName={joinGroupName}
            setJoinGroupName={setJoinGroupName}
            handleJoinGroup={handleJoinGroup}

            leftGroupName={leftGroupName}
            setLeftGroupName={setLeftGroupName}
            handleLeftGroup={handleLeftGroup}

            channels={channels}
            connectedUsername={connectedUsername}
          />
          <div className="chatWindow">
            <Messages
              msgs={msgs}
            />
            <InputMessage
              content={content}
              setContent={setContent}
              receiver={receiver}
              setReceiver={setReceiver}
              handleSendMessage={handleSendMessage}
            />
            <ToastContainer delay={3000} />
          </div>
        </div>
      </div>
    </div>
  );
}

const customCompare = (a, b) => {
  if (a.type < b.type) {
    return -1
  }
  if (a.type > b.type) {
    return 1
  }
  if (a.identifier < b.identifier) {
    return -1
  }
  if (a.identifier > b.identifier) {
    return 1
  }
  return 0
}