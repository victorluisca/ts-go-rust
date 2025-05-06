use std::path::PathBuf;

use anyhow::{Context, Ok, Result, anyhow};

use crate::opts::ProjectorOptions;

#[derive(Debug)]
pub struct Config {
    pub operation: Operation,
    pub pwd: PathBuf,
    pub config: PathBuf,
}

impl TryFrom<ProjectorOptions> for Config {
    type Error = anyhow::Error;

    fn try_from(value: ProjectorOptions) -> Result<Self> {
        let operation = value.arguments.try_into()?;
        let config = get_config(value.config)?;
        let pwd = get_pwd(value.pwd)?;

        return Ok(Config {
            operation,
            config,
            pwd,
        });
    }
}

#[derive(Debug)]
pub enum Operation {
    Print(Option<String>),
    Add(String, String),
    Delete(String),
}

impl TryFrom<Vec<String>> for Operation {
    type Error = anyhow::Error;

    fn try_from(value: Vec<String>) -> Result<Self> {
        let mut value = value;

        if value.len() == 0 {
            return Ok(Operation::Print(None));
        }

        let term = value.get(0).expect("to exist");
        if term == "add" {
            if value.len() != 3 {
                return Err(anyhow!(
                    "operation add expects 2 arguments but got {}",
                    value.len() - 1
                ));
            }

            let mut drain = value.drain(1..=2);

            return Ok(Operation::Add(
                drain.next().expect("to exist"),
                drain.next().expect("to exist"),
            ));
        }

        if term == "del" {
            if value.len() != 2 {
                return Err(anyhow!(
                    "operation delete expects 1 argument but got {}",
                    value.len() - 1
                ));
            }

            let arg = value.pop().expect("to exist");
            return Ok(Operation::Delete(arg));
        }

        if value.len() > 1 {
            return Err(anyhow!(
                "operation print expects 0 or 1 arguments but got {}",
                value.len() - 1
            ));
        }

        let arg = value.pop().expect("to exist");
        return Ok(Operation::Print(Some(arg)));
    }
}

fn get_config(config: Option<PathBuf>) -> Result<PathBuf> {
    if let Some(value) = config {
        return Ok(value);
    }

    let location = std::env::var("XDG_CONFIG_HOME").context("unable to get XDG_CONFIG_HOME")?;
    let mut location = PathBuf::from(location);

    location.push("projector");
    location.push("projector.json");

    return Ok(location);
}

fn get_pwd(pwd: Option<PathBuf>) -> Result<PathBuf> {
    if let Some(pwd) = pwd {
        return Ok(pwd);
    }

    return Ok(std::env::current_dir().context("error getting current directory")?);
}
