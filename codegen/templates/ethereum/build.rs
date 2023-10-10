use anyhow::{Ok, Result};
use regex::Regex;
use substreams_ethereum::Abigen;
use std::fs;
use std::fs::File;
use std::io::Write;

fn main() -> Result<(), anyhow::Error> {
    let file_names = [
        "abi/contract.abi.json",
    ];
    let file_output_names = [
        "src/abi/contract.rs",
    ];

    let mut i = 0;
    for f in file_names {
        let contents = fs::read_to_string(f)
            .expect("Should have been able to read the file");

        // sanitize fields and attributes starting with an underscore
        let regex = Regex::new(r#"("\w+"\s?:\s?")_(\w+")"#).unwrap();
        let sanitized_abi_file = regex.replace_all(contents.as_str(), "${1}u_${2}");

        // do not modify the original abi
        let mut file = File::create("/tmp/contract.abi.json")?;
        file.write_all(sanitized_abi_file.as_bytes())?;

        Abigen::new("Contract", "/tmp/contract.abi.json")?
            .generate()?
            .write_to_file(file_output_names[i])?;

        i = i+1;
    }

    Ok(())
}
