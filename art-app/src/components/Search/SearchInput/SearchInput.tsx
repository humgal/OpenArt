import React from "react";
import { TextInput, View } from "react-native";
import style from "./SearchInput.css";
import { IconButton } from "react-native-paper";
import Icon from "@expo/vector-icons/MaterialIcons";

const SearchInput = () => {
  return (
    <View>
      <TextInput style={style.SearchInput} placeholder="Search Item" />
      <IconButton
        style={style.SearchIcon}
        size={18}
        icon={(props) => <Icon name="search" {...props} />}
      ></IconButton>
      <IconButton
        style={style.SearchVoice}
        size={18}
        icon={(props) => <Icon name="keyboard-voice" {...props} />}
      ></IconButton>
    </View>
  );
};

export default SearchInput;
