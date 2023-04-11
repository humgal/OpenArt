
import React from 'react'
import { View,Text,Button} from 'react-native';


const SearchOption = ()=>{
    return(
        <View style={{margin:5}}>
        <Text> Types</Text>
        <View  style={{flexDirection:'row', flexWrap: 'wrap' }} >
        <Button title="All Item" />
        <Button title = "Game" />
        <Button title = "Video" />
        <Button title = "Animation"  />
        <Button title = "Photography" />
        </View>
        </View>

    )
}

export default SearchOption;