import React, { useState } from "react";
import { View, Text } from "react-native";
import { Button } from "react-native-paper";
import style from "./SearchOption.css";

const SearchOption = () => {
  const [itemsColor, setItemColor] = useState("lightgrey");
  const [gameColor, setGamecolor] = useState("lightgrey");
  const [videoColor, setVideocolor] = useState("lightgrey");
  const [animationColor, setAnimationColor] = useState("lightgrey");
  const [photographyColor, setPhotographyColor] = useState("lightgrey");

  function setDefaultColor() {
    setItemColor("lightgrey");
    setGamecolor("lightgrey");
    setVideocolor("lightgrey");
    setAnimationColor("lightgrey");
    setPhotographyColor("lightgrey"); 
  }

  return (
    <View style={{ margin: 5 }}>
      <Text> Type</Text>
      <View style={{ flexDirection: "row", flexWrap: "wrap" }}>
        <Button
          mode="contained-tonal"
          onPress={() => {
            setDefaultColor()
            setItemColor("mediumpurple");
            
          }}
          style={[style.type, { backgroundColor: itemsColor }]}
        >
          All Items
        </Button>
        <Button
          mode="contained-tonal"
          onPress={() => {
            setDefaultColor()
            setGamecolor("mediumpurple");
           
          }}
          style={[style.type, { backgroundColor: gameColor }]}
        >
          Game
        </Button>
        <Button
          mode="contained-tonal"
          onPress={() => {
            setDefaultColor()
            setVideocolor("mediumpurple");
           
          }}
          style={[style.type, { backgroundColor: videoColor }]}
        >
          Video
        </Button>
        <Button
          mode="contained-tonal"
          onPress={() => {
            setDefaultColor()
            setAnimationColor("mediumpurple");
            
          }}
          style={[style.type, { backgroundColor: animationColor }]}
        >
          Animation
        </Button>

        <Button
          mode="contained-tonal"
          onPress={() => {
            setDefaultColor()
            setPhotographyColor("mediumpurple");
           
          }}
          style={[style.type, { backgroundColor: photographyColor }]}
        >
          Photography
        </Button>
      </View>
    </View>
  );
};

export default SearchOption;
