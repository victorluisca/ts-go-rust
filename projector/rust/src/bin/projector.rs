use anyhow::Result;
use clap::Parser;
use projector::{
    config::{Config, Operation},
    opts::ProjectorOptions,
    projector::Projector,
};

fn main() -> Result<()> {
    let config: Config = ProjectorOptions::parse().try_into()?;
    let mut projector = Projector::from_config(config.config, config.pwd);

    match config.operation {
        Operation::Print(None) => {
            let value = projector.get_value_all();
            let value = serde_json::to_string(&value)?;

            print!("{}", value);
        }
        Operation::Print(Some(k)) => {
            if let Some(x) = projector.get_value(&k) {
                println!("{}", x);
            }
        }
        Operation::Add(k, v) => {
            projector.set_value(k, v);
            projector.save()?;
        }
        Operation::Delete(k) => {
            projector.delete_value(&k);
            projector.save()?;
        }
    }

    return Ok(());
}
