import { Avatar, Button, Card } from "react-native-paper";
import { User } from "../../../gql/types";
import { Image, View, Text, StyleSheet } from "react-native";
import Icon from "@expo/vector-icons/Ionicons";
import style from "./creatorcard.css";

const CreatorCard = (props: User) => {
  return (
    <Card style={{ padding: 5, margin: 8 }}>
      <Card.Cover style={{ width: "100%" }} source={{ uri: props.img }} />
      <View>
        <Avatar.Image style={style.avatar} source={require("./2.png")} />
      </View>
      <View>
        <Text style={style.username}>{props.username}</Text>
        <Text style={style.bio}>{props.bio}</Text>
        <View style={{height:30}}>
          <Text style={style.followerNum}>  
          <Text style={{fontSize:20}}>
          {props.followerNum }
          </Text>
          
          Followers</Text>
          <Text style={style.followButton} onPress={() => {}}>
            Follow
          </Text>
        </View>
      </View>
    </Card>
  );
};

export default CreatorCard;
