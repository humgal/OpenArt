import { Button } from "react-native-paper";
import { User } from "../../../gql/types";
import {Image, View,Text,StyleSheet} from "react-native";
import { Icon } from "@react-native-material/core";

const CreatorCard = (props:User) =>{
    return(
        <View>
            <Image style={styles.img} source={require(String(props.img))} /> 
            <Image style={styles.avatar}  source={require(String(props.avatar))}/>
            <Text style={styles.username} >{props.username}</Text>
            <Text style={styles.description} >{props.bio}</Text>
            <Text style={styles.followerNum} > {props.followerNum}</Text>
            <Button style={styles.followButton} icon={props => <Icon name="md-person-outline" {...props} />}
            onPress={() => {
              
            }}>Follow</Button>
        </View>
        
    )
}

const styles = StyleSheet.create({
    img:{

    },
    avatar:{

    },
    username:{

    },
    description:{

    },
    followerNum:{

    },
    followButton:{

    }
})

export default CreatorCard;