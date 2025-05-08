import getOptions from "./opts";
import getConfig, { Operation } from "./config";
import Projector from "./projector";

const options = getOptions();
const config = getConfig(options);
const projector = Projector.fromConfig(config);

if (config.operation === Operation.Print) {
    if (config.arguments.length === 0) {
        console.log(JSON.stringify(projector.getValueAll()));
    } else {
        const value = projector.getValue(config.arguments[0]);
        if (value) {
            console.log(value);
        }
    }
}

if (config.operation === Operation.Add) {
    projector.setValue(config.arguments[0], config.arguments[1]);
    projector.save();
}

if (config.operation === Operation.Delete) {
    projector.deleteValue(config.arguments[0]);
    projector.save();
}
