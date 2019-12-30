import React from "react";

class InputCheckbox extends React.Component {
  constructor(props, context) {
    super(props, context);
    this.selected = {};
  }

  render() {
    const { boxSelect, onChangeSelect } = this.props;
    return (
        <div className="form-group">
            {Object.keys(boxSelect).map((skill) =>  {
                return (
                    <div className="form-check" key={skill}>
                        <label>Skill </label> 
                        {skill}
                        <input 
                            type="checkbox" 
                            value={skill}
                            checked={boxSelect[skill]}
                            onChange={e => onChangeSelect(skill)}
                        />{" "}
                    </div>    
                )} )}
        </div>
    );
  }
}

export default InputCheckbox;

