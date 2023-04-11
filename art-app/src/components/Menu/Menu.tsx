import { View, Text,Alert,Button } from "react-native";

const Menu = () => {
  return (
    <View style={{ flex: 1 }}>
      <View style={{ flex: 5}}>
        <View  style={{ flex: 2}}></View>
        <Text style={{ flex: 1,alignSelf:"center" , fontSize:30}}>About OpenArt</Text>
        <Text style={{ flex: 1 ,alignSelf:"center",fontSize:30}}>Blog</Text>
        <Text style={{ flex: 1 ,alignSelf:"center",fontSize:30}}>Help</Text>
        <Text style={{ flex: 1 ,alignSelf:"center",fontSize:30}}>Contact</Text>
      </View>
      <View style={{ flex: 5 }}>
        <View  style={{ flex: 3}}></View>
        <View  style={{ flex: 2, width:'70%',alignSelf:'center'}}>
        <Button title="Connect Wallet" onPress={() => {}} />
        </View>
            
      </View>
    </View>
  );
};

export default Menu;
